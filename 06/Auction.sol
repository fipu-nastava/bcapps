pragma solidity ^0.5.0;

contract StateMachine {

    enum Stages {
        AcceptingBlindedBids,
        RevealBids,
        Korak1,
        Korak2,
        Korak3
    }

    struct Bid {
        address bidder;
        uint amount;
    }

    Bid[] bids;

    // Trenutno stanje
    Stages public stage = Stages.AcceptingBlindedBids;

    uint public creationTime = now; // ovo se evaluira pri deploymentu

    modifier atStage(Stages _stage) {
        require(stage == _stage, "Poziv nije moguć u ovom trenutku.");
        _;
    }

    modifier minStage(Stages _stage) {
        require(stage >= _stage, "Poziv nije moguć u ovom trenutku.");
        _;
    }

    function nextStage() internal {
        // pretvori enum u integer, uvećaj i pretvori natrag
        stage = Stages(uint(stage) + 1);
    }

    // Modifier koji pazi da smo u ispravnoj fazi ovisno o vremenu
    // Pošto se kod ne može izrvšavati bez transakcija, ovo će se izrvšiti prije transakcije
    modifier timedTransitions() {
        // 10 dana traje prvi krug
        if (stage == Stages.AcceptingBlindedBids && now >= creationTime + 5 minutes) {
            nextStage();
        }
        // zatim 2 dana traje sljedeća faza 
        if (stage == Stages.RevealBids && now >= creationTime + 10 minutes) {
            nextStage();
        }
        // The other stages transition by transaction
        _;
    }

    // Paziti na poredak modifiera, prvo želimo da se uveća faza (ako treba)
    function bid(uint amount) public payable timedTransitions atStage(Stages.AcceptingBlindedBids)
    {
        Bid b = Bid(msg.sender, amount);
        bids.push(b);
    }

    function reveal() public timedTransitions atStage(Stages.RevealBids)
    {
    }

    function getBidCount() public view minStage(Stages.RevealBids) returns (uint) {
        return bids.length;
    } 

    function getBidAmountAt(uint index) public view minStage(Stages.RevealBids) returns (uint)
    {
        require(index >= 0 && index < bids.length);
        return bids[index].amount;
    }

    // Uvećaj fazu, ali nakon izrvšavanje funkcije/transakcije
    modifier transitionNext()
    {
        _;
        nextStage();
    }

    function korak1() public timedTransitions atStage(Stages.Korak1) transitionNext
    {
    }

    function korak2() public timedTransitions atStage(Stages.Korak2) transitionNext
    {
    }

    function korak3() public timedTransitions atStage(Stages.Korak3)
    {
    }
}