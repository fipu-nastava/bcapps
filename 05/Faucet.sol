pragma solidity ^0.5.1;

// Primjer prvog ugovora - faucet
contract Faucet {

    // Pokloni ether svima koji pitaju
    function withdraw(uint withdraw_amount) public {

        // Limit koliko se može zatražiti
        // Baca Exception ako nije OK
        require(withdraw_amount <= 100000000000000000);

        // Slanje iznosa na adresu koja je zatražila
        msg.sender.transfer(withdraw_amount);
    }

    // Primi bili koju donaciju, defaultna funkcija
    function () external payable {}

}
