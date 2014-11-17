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

  function VReel(symbols,game,x,y,mask){
    this.reel=game.add.group();

    this.firstHalf=game.add.group();
    _.each(symbols,function(e,i,l){this.firstHalf.create(0,height*i,e);},this);
    this.reel.add(this.firstHalf);

    this.secondHalf=game.add.group();
    _.each(symbols,function(e,i,l){this.secondHalf.create(0,height*i,e);},this);
    this.reel.add(this.secondHalf);
    
    this.firstHalf.y= -height*symbols.length;

    this.reel.x=x;
    this.reel.y=y;
    this.reel.mask=mask;
    this.total=symbols.length*2;
  }

  VReel.prototype.changeSymbols=function(newSymbols){
    this.firstHalf.removeAll(true,true);
    this.secondHalf.removeAll(true,true);
    _.each(newSymbols,function(e,i,l){this.firstHalf.create(0,height*i,e);},this);
    _.each(newSymbols,function(e,i,l){this.secondHalf.create(0,height*i,e);},this);
    this.secondHalf.y = 0;
    this.firstHalf.y= -height*newSymbols.length;
    this.total=newSymbols.length*2;
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
      var p=jQuery.getJSON('/game/underwater/spin/3/4');
      var observable=Rx.Observable.fromPromise(p);
      observable.delay(3000).subscribe(function(data){
        console.log("in rxjs");
        console.log(data);
        var symbols=['symbol_Nine','symbol_Ten','symbol_Jack'];
        //
        that.spinning=false;
        that.vreel.changeSymbols(symbols);
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
      
      
      var reel0=['symbol_Nine','symbol_Ace','symbol_Nemo','symbol_King','symbol_Nemo'];
      this.vreel=new VReel(reel0,this,reelsleft,reelstop,mask);  
      
    },

    update: function () {
      if(this.spinning){
        if(this.vreel.firstHalf.y>boxHeight){
          this.vreel.firstHalf.y -= this.vreel.total*height;
        }
        if(this.vreel.secondHalf.y>boxHeight){
          this.vreel.secondHalf.y -= this.vreel.total*height;
        }
        this.vreel.firstHalf.y+=15;
        this.vreel.secondHalf.y+=15;
      }
    }
  };

  window['underwater'] = window['underwater'] || {};
  window['underwater'].Game = Game;

}());
