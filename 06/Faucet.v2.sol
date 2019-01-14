pragma solidity ^0.5.1;

// Primjer prvog ugovora - faucet
contract Faucet {

    address owner;

    constructor() {
        this.owner = msg.sender;
    }

    modifier onlyOwner {
        require(msg.sender == this.owner, "Nisi vlasnik.");
        _;
    }

    // Pokloni ether svima koji pitaju
    function withdraw(uint withdraw_amount) public {

        // Limit koliko se može zatražiti
        // Baca Exception ako nije OK
        require(withdraw_amount <= 0.1 ether);

        // Slanje iznosa na adresu koja je zatražila
        msg.sender.transfer(withdraw_amount);
    }

    // Primi bili koju donaciju, defaultna funkcija
    function () external payable {}

    // Metoda koja će brisati ugovor, može se nazvati kako god (npr. krumpir)
    function destroy() public onlyOwner {

        // Poništi ugovor i prebaci preostala sredstva na adresu vlasnika 
        selfdestruct(owner);
    }
}
