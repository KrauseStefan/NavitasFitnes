package ExportService

import (
	"net/http"
	"testing"

	"appengine"
	"appengine/datastore"

	"NavitasFitness/TestHelper"
	"NavitasFitness/User/Dao"
)

var assert = TestHelper.Assert

func mockoutGetAllUsers(keys []*datastore.Key, users []UserDao.UserDTO, e error) *TestHelper.Spy {
	spy := new(TestHelper.Spy)
	userDao_GetAllUsers = func(ctx appengine.Context) ([]*datastore.Key, []UserDao.UserDTO, error) {
		_ctx := &ctx
		spy.RegisterCall()
		spy.RegisterArg1(_ctx)
		return keys, users, e
	}
	return spy
}

func TestShouldConfigureHeaderForDownload(t *testing.T) {
	header := make(http.Header)
	configureHeaderForFileDownload(&header, "test.file")

	assert(t, header.Get("Content-Disposition") == "attachment; filename=test.file")
	assert(t, header.Get("Content-Type") == "application/vnd.ms-excel")
}

func TestShouldConfigureHeaderForNoCache(t *testing.T) {
	header := make(http.Header)
	configureHeaderForFileDownload(&header, "test.file")

	assert(t, header.Get("Cache-Control") == "no-cache, no-store, must-revalidate")
	assert(t, header.Get("Pragma") == "no-cache")
	assert(t, header.Get("Expires") == "0")
}

//func TestShouldGetTransactionsFromDataStore(t *testing.T) {
//	ctx := new(TestHelper.ContextMock)
//
//	spy := mockoutGetAllUsers(make([]*datastore.Key, 0, 0), make([]UserDao.UserDTO, 0, 0), nil)
//
//	getTransactionList(ctx)
//
//	assert(t, spy.CallCount()).Equals(1)
//	assert(t, spy.GetLatestArg1()).Equals(ctx)
//}

func TestExportXsltHandler(t *testing.T) {

	ctx := new(TestHelper.ContextMock)

	exportXslt(ctx)

}
