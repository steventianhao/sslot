(function () {
  'use strict';

  function Boot() {}

  Boot.prototype = {
    
    preload: function () {
      this.load.image("loadingbg",'assets/01_intro.jpg')
      this.load.image('preloader', 'assets/preloader.gif');
    },

    create: function () {
      this.game.input.maxPointers = 1;

      if (this.game.device.desktop) {
        this.game.scale.pageAlignHorizontally = true;
      } else {
        this.game.scaleMode = Phaser.ScaleManager.SHOW_ALL;
        this.game.scale.minWidth =  1024;
        this.game.scale.minHeight = 768;
        this.game.scale.maxWidth = 1024;
        this.game.scale.maxHeight = 768;
        this.game.scale.forceLandscape = true;
        this.game.scale.pageAlignHorizontally = true;
        this.game.scale.setScreenSize(true);
      }
      this.game.state.start('preloader');
    }
  };

  window['underwater'] = window['underwater'] || {};
  window['underwater'].Boot = Boot;

}());

