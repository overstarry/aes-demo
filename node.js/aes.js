var crypto = require("crypto");

function aesEncode(data, secret,iv) {
    const cipher = crypto.createCipheriv('aes-128-cbc', secret, iv);
    var crypted = cipher.update(data, "binary", "hex");
    crypted += cipher.final('hex');
    return crypted
}

function aesDecode(data, secret,iv) {
    const cipher = crypto.createDecipheriv('aes-128-cbc', secret, iv);
    var crypted = cipher.update(data, "hex", "binary");
    crypted += cipher.final('binary');
    return crypted
}

const date_crypto = aesEncode("i am test data", "abcdefghabcdefgh","1234567812345678")
console.log(date_crypto)
console.log(aesDecode(date_crypto, "abcdefghabcdefgh","1234567812345678"))