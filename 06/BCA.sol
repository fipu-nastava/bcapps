pragma solidity ^0.5.0;
 
contract BCA {
    uint256 kolicinaNovca = 1000000;
    
    // Stanje novaca po računima
    mapping(address => uint256) stanje;
 
    // Konstruktor
    constructor() public {
        stanje[msg.sender] = kolicinaNovca;
    }
 
    // Koliko ima pojedini račun?
    function balanceOf(address _adresa) public view returns (uint256 balance) {
        return stanje[_adresa];
    }
 
    // Transfer novaca
    function transfer(address _to, uint256 _amount) public returns (bool success) {
        if (stanje[msg.sender] >= _amount 
            && _amount > 0
            && stanje[_to] + _amount > stanje[_to]) {
            stanje[msg.sender] -= _amount;
            stanje[_to] += _amount;
            return true;
        } else {
            return false;
        }
    }
}