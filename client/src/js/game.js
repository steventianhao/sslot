(function() {
  'use strict';

  var reelsleft=139;
  var reelstop=110;



  var boxWidth=736;
  var boxHeight=409;

  var wGap=8;
  var hGap=2;

  var width=141+wGap;

  var height=135+hGap;

  var reelsbottom=110+height*3;

  var spin=0;

  function Game() {
    this.player = null;
    this.spinning=false;
    this.reel0 =null;
    this.reel1=null;
    this.reel2=null;
    this.reel3=null;

    this.symbols=['symbol_Nine','symbol_Ace','symbol_Nemo','symbol_Mermaid','symbol_Ten','symbol_Starfish','symbol_King','symbol_Queen','symbol_Clam','symbol_Jack','symbol_Shark','symbol_Green','symbol_Octopus']
  }

  function Reel(hits,pads){
    this.hits = hits;
    this.pads = pads;
    this.total = hits.length+pads.length;
  }

  function VReel(reel,game){
    var g=game.add.group();
    var g1=game.add.group();
    _.each(reel.pads,function(e,i,l){g1.create(0,height*i,e);});
    g.add(g1);

    var g0=game.add.group();
    _.each(reel.hits,function(e,i,l){g0.create(0,height*i,e);});
    g.add(g0);
    g0.y= g0.y-height*(reel.hits.length);
    
    return {reel:g,hits:g0,pads:g1,total:reel.total};
  }


  function gen(allsymbols, expected){
    var total=0;
    var result=[];
    while(total<expected){
      var idx= _.random(allsymbols.length-1);
      result.push(allsymbols[idx]);
      total=total+1;
    }
    return result;
  }

  function spinPressed(){
    this.spinning=true;    
    if(this.spinning){
      var that =this;
      var p=jQuery.getJSON('http://localhost:5555/game/underwater/spin/3/4');
      var observable=Rx.Observable.fromPromise(p);
      observable.delay(10000).subscribe(function(data){
        console.log("in rxjs");
        console.log(data);
        that.spinning=false;
      });
    }
  }

  Game.prototype = {

    preload: function(){
      this.stage.disableVisibilityChange = true;
    },

    create: function(){
      
      this.player = this.add.sprite(0, 0, 'background');
      this.add.sprite(110,160,"numberleft");
      this.add.sprite(reelsleft,reelstop,"maingamebg");
      this.add.sprite(139+736,160,"numberright");

      var mask = this.add.graphics(0, 0);
      mask.beginFill(0xffffff);
      mask.drawRect(reelsleft, reelstop, boxWidth,boxHeight);
      
      this.bSpinButton=this.add.button(745,583,"spin", spinPressed, this, "spin01.png","spin02.png","spin03.png");
      
      
      var reel=new Reel(['symbol_Ace','symbol_King','symbol_Queen'],['symbol_Nine','symbol_Ace','symbol_Nemo','symbol_Nemo','symbol_Nemo']);
      this.vreel=new VReel(reel,this);
      var g=this.vreel.reel;

      g.x=reelsleft;
      g.y=reelstop;

      
    
      
      g.mask=mask;
      
    },

    update: function () {
      if(this.spinning){
        if(this.vreel.hits.y>boxHeight){
          this.vreel.hits.y -= this.vreel.total*height;
        }
        if(this.vreel.pads.y>boxHeight){
          this.vreel.pads.y -= this.vreel.total*height;
        }
        this.vreel.hits.y+=15;
        this.vreel.pads.y+=15;
      }
    }
  };

  window['underwater'] = window['underwater'] || {};
  window['underwater'].Game = Game;

}());
