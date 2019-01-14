async function init() {
    $out = $("#out");
    $out.val("");

    var rpc = $("#web3").val();

    if (rpc) {
        window.web3 = new Web3(new Web3.providers.HttpProvider(rpc));
    }
    else {
        // MetaMask
        rpc = "MetaMask";
        window.web3 = new Web3(window.ethereum);
        await window.ethereum.enable();
    }

    var accounts = []
    try {
        accounts = web3.eth.accounts
        console.log(accounts);
        $("#out").val([`Successfully connected to ${rpc}!`, "Accounts present:", ""].concat(web3.eth.accounts).join("\n"))
    }
    catch(e) {
        console.log(e);
        $("#out").val(e);
    }
}

function transaction() {
    // src = web3.eth.accounts[0];
    // dst = web3.eth.accounts[1];
    var src = $("#from").val();
    var dst = $("#to").val();
    var amount = $("#amount").val();

    web3.eth.sendTransaction({
        from: src,
        to: dst,
        value: web3.toWei(amount, "ether")
    }, handleResult)
}

function handleResult(err, txid) {
    if (!err) {
        alert("Transaction created with txid: " + txid);
        console.log(err, txid);
        return
    }
    else {
        $("#out").val(err);
    }
}

function deploy() {
    var src = web3.eth.accounts[0];
    // bytecode_data = "0x608060405234801561001057600080fd5b5060fd8061001f6000396000f3fe6080604052600436106039576000357c0100000000000000000000000000000000000000000000000000000000900480632e1a7d4d14603b575b005b348015604657600080fd5b50607060048036036020811015605b57600080fd5b81019080803590602001909291905050506072565b005b67016345785d8a00008111151515608857600080fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801560cd573d6000803e3d6000fd5b505056fea165627a7a723058204cfa2dc3b895145b483970ca217335a22c50693fb547548d56570e3af5f04b0f0029"
    var bytecode_data = $("#bytecode").val()

    web3.eth.sendTransaction({
        from: src,
        gas: 3000000,
        data: bytecode_data
    }, handleResult)
}

function clearAll() {
    var fields = ["out", "from", "to", "web3", "code", "txid", "amount"]

    for(var f of fields) {
        $("#"+f).val("");
    }
}

function getTxStatus() {
    var txid = $("#txid").val();

    web3.eth.getTransactionReceipt(txid, (err, receipt) => {
        console.log(receipt);
        $("#out").val(JSON.stringify(receipt))
    })
}
