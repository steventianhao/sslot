window.onload = function () {
  'use strict';

  var game
    , ns = window['underwater'];

  game = new Phaser.Game(1024, 768, Phaser.AUTO, 'underwater-game');
  game.state.add('boot', ns.Boot);
  game.state.add('preloader', ns.Preloader);
  game.state.add('menu', ns.Menu);
  game.state.add('game', ns.Game);

  game.state.start('boot');
};
