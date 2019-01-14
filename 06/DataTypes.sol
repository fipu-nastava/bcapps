pragma solidity ^0.5.0;

contract DataTypes {
    
    bool public b;
    uint public i;
    int public j;
    address public adresa;
    string public nazivUgovora;
    
    enum Stanje {Prvo, Drugo, Trece} 
    struct Osoba {string Ime; string Prezime; }
    
    Stanje public stanjeUgovora;
    Osoba public vlasnikUgovora;
    
    uint[] public polje_intova;
    
    mapping (uint => uint) public stanjeNovaca;
    
    function usporedba() public view returns(bool) {
        return keccak256("DataTypes") == keccak256(bytes(nazivUgovora));
    }
    
    function duljina() public view returns(uint) {
        return polje_intova.length;
    }
    
    constructor() public {
        b = true;
        i = 1;
        i -= 2;
        j = 1;
        j -= 2;
        polje_intova.push(1);
        polje_intova.push(2);
        nazivUgovora = "DataTypes";
        adresa = address(this);
        stanjeUgovora = Stanje.Prvo;
        stanjeUgovora = Stanje(uint8(stanjeUgovora)+1);
        stanjeNovaca[0] = 100;
        stanjeNovaca[1] = 101;
        vlasnikUgovora.Ime = "Nikola";
        vlasnikUgovora.Prezime = "Tankovic";
    }
}