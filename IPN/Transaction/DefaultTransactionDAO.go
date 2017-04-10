package TransactionDao

import (
	"errors"
	"fmt"
	"time"

	"appengine"
	"appengine/datastore"

	"User/Dao"
)

var userDao = UserDao.GetInstance()

type DefaultTransactionDao struct{}

var defaultTransactionDao = DefaultTransactionDao{}

func GetInstance() TransactionDao {
	return &defaultTransactionDao
}

var (
	TxnDuplicateTxnMsg = errors.New("Doublicate message recived, this is likely not a programming error")
	//txnUnableToVerify = errors.New("Unable to verify message")
)

func (t *DefaultTransactionDao) UpdateIpnMessage(ctx appengine.Context, ipnTxn *TransactionMsgDTO) error {

	key := ipnTxn.GetDataStoreKey(ctx)

	//Make sure indexed fields are updated
	ipnTxn.parseMessage()

	if _, err := datastore.Put(ctx, key, ipnTxn.dsDto); err != nil {
		return err
	}

	return nil
}

func (t *DefaultTransactionDao) PersistNewIpnMessage(ctx appengine.Context, ipnTxn *TransactionMsgDTO, userKey *datastore.Key) error {

	var newKey *datastore.Key

	if ipnTxn.hasKey() {
		return errors.New("ipnTxn has already been persisted, use update function ínstead")
	}

	if userKey == nil {
		newKey = datastore.NewIncompleteKey(ctx, TXN_KIND, txnCollectionParentKey(ctx))
	} else {
		newKey = datastore.NewIncompleteKey(ctx, TXN_KIND, userKey)
	}

	//Make sure indexed fields are updated
	ipnTxn.parseMessage()
	if _, err := datastore.Put(ctx, newKey, &ipnTxn.dsDto); err != nil {
		return err
	}

	return nil
}

func (t *DefaultTransactionDao) GetTransaction(ctx appengine.Context, txnId string) (*TransactionMsgDTO, error) {
	q := datastore.NewQuery(TXN_KIND).
		Filter("TxnId=", txnId).
		Limit(1)

	txnDtoList := make([]transactionMsgDsDTO, 0, 1)

	keys, err := q.GetAll(ctx, &txnDtoList)
	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return nil, nil
	}

	return NewTransactionMsgDTOFromDs(txnDtoList[0], keys[0]), nil
}

func (t *DefaultTransactionDao) GetTransactionsByUser(ctx appengine.Context, parentUserKey *datastore.Key) ([]*TransactionMsgDTO, error) {

	q := datastore.NewQuery(TXN_KIND).
		Ancestor(parentUserKey).
		Order("PaymentDate")

	count, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	txnDsDtoList := make([]transactionMsgDsDTO, 0, count)

	keys, err := q.GetAll(ctx, &txnDsDtoList)
	if err != nil {
		return nil, err
	}

	return NewTransactionMsgDTOList(txnDsDtoList, keys), nil
}

func (t *DefaultTransactionDao) GetCurrentTransactionsAfter(ctx appengine.Context, userKey *datastore.Key, date time.Time) ([]*TransactionMsgDTO, error) {

	q := datastore.NewQuery(TXN_KIND).
		Ancestor(userKey).
		Filter("PaymentDate>=", date)

	count, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	txnDsDtoList := make([]transactionMsgDsDTO, 0, count)

	keys, err := q.GetAll(ctx, &txnDsDtoList)
	if err != nil {
		return nil, err
	}

	if count > 1 {
		ctx.Criticalf(fmt.Sprintf("User has multiple (%d) active subscriptions, key: %s", count, userKey.String()))
	}

	return NewTransactionMsgDTOList(txnDsDtoList, keys), nil
}
