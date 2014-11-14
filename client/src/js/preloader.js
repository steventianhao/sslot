(function() {
  'use strict';

  function Preloader() {
    this.asset = null;
    this.ready = false;
  }

  Preloader.prototype = {

    preload: function () {
      this.add.sprite(0,0,'loadingbg');
      this.asset = this.add.sprite(1024/2, 768/2, 'preloader');
      this.asset.anchor.setTo(0.5, 0.5);
      this.load.onLoadComplete.addOnce(this.onLoadComplete, this);
      this.load.setPreloadSprite(this.asset);
      
      this.load.image('symbol_Nine','assets/symbols/9_00000.png');
      this.load.image('symbol_Ten','assets/symbols/10_00000.png');
      this.load.image('symbol_Jack','assets/symbols/J_00000.png');
      this.load.image('symbol_Queen','assets/symbols/Q_00000.png');
      this.load.image('symbol_King','assets/symbols/K_00000.png');
      this.load.image('symbol_Ace','assets/symbols/A_00000.png');
      this.load.image('symbol_Clam','assets/symbols/贝壳_00000.png');
      this.load.image('symbol_Starfish','assets/symbols/海星_00000.png');
      this.load.image('symbol_Nemo','assets/symbols/小丑鱼_00000.png');
      this.load.image('symbol_Green','assets/symbols/海龟_00000.png');
      this.load.image('symbol_Octopus','assets/symbols/章鱼_00000.png');
      this.load.image('symbol_Mermaid','assets/symbols/美人鱼_00000.png');
      this.load.image('symbol_Shark','assets/symbols/鲨鱼01_00000.png');

      this.load.image('background','assets/background.jpg');
      this.load.image('freegamebg','assets/free games02 water BG.jpg');
      this.load.image('maingamebg','assets/NORMAL games water BG.jpg');

      this.load.image('numberleft','assets/number_left.png');
      this.load.image('numberright','assets/number_right.png');

      this.load.atlas('spin','assets/buttons/spinatlas.png',"assets/buttons/spinatlas.json");
    },

    create: function () {
      this.asset.cropEnabled = false;
    },

    update: function () {
      if (!!this.ready) {
        this.game.state.start('game');
      }
    },

    onLoadComplete: function () {
        this.ready = true;  
    }
  };

  window['underwater'] = window['underwater'] || {};
  window['underwater'].Preloader = Preloader;

}());
