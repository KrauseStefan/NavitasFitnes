import { module } from 'angular';

import { MainPageService } from './PageComponents/MainPage/MainPageService';
import { UserService } from './PageComponents/UserService';

import { AppComponent } from './AppComponent';
import { CkEditor } from './Components/CkEditor/CkEditor';
import { AdminPageComponent, adminRouterState } from './PageComponents/AdminPage/AdminPageComponent';
import { MainPageComponent, mainPageRouterState } from './PageComponents/MainPage/MainPage';
import { nfResetOnChange } from './PageComponents/RegistrationForm/nfResetOnChange';
import { nfShouldEqual } from './PageComponents/RegistrationForm/nfShouldEqual';
import { UserStatusComponent, statusRouterState } from './PageComponents/UserStatus/UserStatus';

import { DefaultErrorHandlingModule } from './DefaultErrorHandling';

import 'angular-animate';
import 'angular-aria';
import 'angular-cookies';
import 'angular-material';
import 'angular-messages';
import 'angular-ui-grid';
import 'angular-ui-router';

import ngMat = ng.material;

export const NavitasFitnessModule = module('NavitasFitness', [
  'ngMaterial', 'ui.router', 'ngCookies', 'ngMessages', 'ui.grid',
  DefaultErrorHandlingModule.name,
])
  .config(($mdThemingProvider: ngMat.IThemingProvider) => {
    $mdThemingProvider.theme('default')
      .primaryPalette('blue')
      .accentPalette('orange');
  })

  .config((
    $stateProvider: ng.ui.IStateProvider,
    $urlRouterProvider: ng.ui.IUrlRouterProvider,
    $locationProvider: ng.ILocationProvider,
    $sceDelegateProvider: ng.ISCEDelegateProvider
  ) => {
    $locationProvider.html5Mode(true);
    $urlRouterProvider.otherwise('/main-page/');

    $stateProvider
      .state('Admin', adminRouterState)
      .state('MainPage', mainPageRouterState)
      .state('Status', statusRouterState);

    $sceDelegateProvider.resourceUrlWhitelist([
      // Allow same origin resource loads.
      'self',
      'http://localhost:8081/processPayment',
      'https://www.sandbox.paypal.com/cgi-bin/webscr',
      'https://www.paypal.com/cgi-bin/webscr',
    ]);

  })

  .service('userService', UserService)
  .service('mainPageService', MainPageService)

  .component('ckEditor', CkEditor)
  .component('mainPage', MainPageComponent)
  .component('userStatus', UserStatusComponent)
  .component('adminPage', AdminPageComponent)
  .component('appComponent', AppComponent)

  .directive(nfResetOnChange.name, nfResetOnChange.factory)
  .directive(nfShouldEqual.name, nfShouldEqual.factory);
