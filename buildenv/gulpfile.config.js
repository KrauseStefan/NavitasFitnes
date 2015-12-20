'use strict';
function GulpConfig() {
  this.outputPath = './app-engine/webapp';

  this.outputLibs = this.outputPath + '/libs';

  this.source = './websrc';
  this.sourceApp = this.source + '/app';

  this.allJavaScript = [this.source + '/js/**/*.js'];
  this.allTypeScript = this.sourceApp + '/**/*.ts';

  this.views = this.sourceApp + '/**/*.jade';

  this.styles = this.source + '/styles/**/*.scss';

  this.typings = this.source + '/typings/';
  this.libraryTypeScriptDefinitions = this.typings + '**/*.ts';

  this.libaryFolders = [
    '/node_modules',
    '/bower_components'
  ].map((lib) => this.source + lib);
}
module.exports = new GulpConfig();