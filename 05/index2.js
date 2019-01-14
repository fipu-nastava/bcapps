function init() {
    // staviti kod za spajanje na RPC 

    var adresa = $("#web3").val();

    // spajanje na RPC

    var provider = new Web3.providers.HttpProvider(adresa);
    window.web3 = new Web3(provider);

    var $out = $("#out")

    var accounts = web3.eth.accounts;
    var ispis = "Uspješno spojeno! \nRačuni: \n\n";
    ispis += accounts.join("\n");

    $out.val(ispis);
}

function callbackFn(err, info) {
    console.log("Prošlo je!");
    console.log("Greška: ", err);
    console.log("Informacije: ", info);

    var $out = $("#out");

    if (!err) {
        $out.val("Transakcija poslana. TxID:\n" + info);
    }
    else {
        $out.val("Transakcija neuspješna. Razlog:\n" + err);
    }
}

function deploy() {
    var bytecode = $("#bytecode").val();
    console.log(bytecode)

    var info = {
        from: web3.eth.accounts[0], // prvi iz liste
        // to: "0x0", // ugovor se kreira slanjem na nultu adresu
        // value: nema
        data: bytecode,
        gas: 300000
    }

    web3.eth.sendTransaction(info, (err, txid) => {
        $out = $("#out");
        console.log(err, txid);
        $out.val("Txid: " + txid);
    });
}

function getTxStatus() {
    var txid = $("#txid").val();

    web3.eth.getTransactionReceipt(txid, (err, info) => {
        console.log(err, info);
        var $out = $("#out");
        $out.val(JSON.stringify(info))
    });
}

function transaction() {
    var src = $("#from").val();
    var dst = $("#to").val();
    var amount = $("#amount").val();

    var weiAmount = web3.toWei(amount, "ether");
    console.log("Wei iznos:", weiAmount);

    var info = {
        from: src,
        to: dst,
        value: weiAmount
    }
    web3.eth.sendTransaction(info, callbackFn);
    console.log("Je li prošlo?");
}