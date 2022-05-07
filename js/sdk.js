const axios = require('axios');

// noinspection DuplicatedCode
const easypay = {
    token: null,
    config: {
        headers: {
            "User-Agent": "EasyPaySDK/js1.0",
            "Content-Type": "application/json"
        }
    },
    init(appId, key) {
        this.token = appId + ":" + key
    },
    _url(method, queryString) {
        let extra = "";
        if (queryString) {
            extra = "?" + queryString
        }
        let host = "pay.easypaybot.com";
        const debug = typeof process.env.EASYPAY_DEBUG === "string" && process.env.EASYPAY_DEBUG !== "";
        if (debug) {
            process.stderr.write("WARNING: DEBUG SERVER USED!!!\n");
            host = 'pay.test.tgbot.link';
        }
        return "https://" + host + "/api/merchant/" + this.token + "/" + method + extra
    },
    _mustOK(resp) {
        let ok = resp && resp.status === 200 && resp.statusText === "OK";
        if (!ok) {
            throw new Error("internal server error")
        }
    },
    async me() {
        const resp = await axios.get(this._url("me"), this.config);
        this._mustOK(resp);
        if (resp.data.code !== 200 || resp.data.status !== "ok") {
            throw new Error("#" + resp.data.code + "(" + resp.data.status + ":" + resp.data.message + ")");
        }
        return resp.data.result;
    },
    async paylink(params) {
        const resp = await axios.post(this._url("paylink"), params, this.config);
        this._mustOK(resp);
        if (resp.data.code !== 200 || resp.data.status !== "ok") {
            throw new Error("#" + resp.data.code + "(" + resp.data.status + ":" + resp.data.message + ")");
        }
        return resp.data.result.items;
    },
    async trans(type, txid) {
        const resp = await axios.get(this._url("trans", "type=" + type + "&txid=" + txid), this.config);
        this._mustOK(resp);
        if (resp.data.code !== 200 || resp.data.status !== "ok") {
            throw new Error("#" + resp.data.code + "(" + resp.data.status + ":" + resp.data.message + ")");
        }
        return resp.data.result;
    },
    async order(id) {
        const resp = await axios.get(this._url("order", "order_id=" + id), this.config);
        this._mustOK(resp);
        if (resp.data.code !== 200 || resp.data.status !== "ok") {
            throw new Error("#" + resp.data.code + "(" + resp.data.status + ":" + resp.data.message + ")");
        }
        return resp.data.result;
    },
    async orderlist(page, pageSize) {
        const resp = await axios.get(this._url("orderlist"), "page=" + page + "&page_size=" + pageSize, this.config);
        this._mustOK(resp);
        if (resp.data.code !== 200 || resp.data.status !== "ok") {
            throw new Error("#" + resp.data.code + "(" + resp.data.status + ":" + resp.data.message + ")");
        }
        return resp.data.result.items;
    },
    async transfer(params) {
        const resp = await axios.post(this._url("transfer"), params, this.config);
        this._mustOK(resp);
        if (resp.data.code !== 200 || resp.data.status !== "ok") {
            throw new Error("#" + resp.data.code + "(" + resp.data.status + ":" + resp.data.message + ")");
        }
        return resp.data.result;
    },
    async deduct(params) {
        const resp = await axios.post(this._url("deduct"), params, this.config);
        this._mustOK(resp);
        if (resp.data.code !== 200 || resp.data.status !== "ok") {
            throw new Error("#" + resp.data.code + "(" + resp.data.status + ":" + resp.data.message + ")");
        }
        return resp.data.result;
    },
    async invitecode(code, tgUserId) {
        const resp = await axios.get(this._url("invitecode", "code=" + code + "&tg_user_id=" + tgUserId), this.config);
        this._mustOK(resp);
        if (resp.data.code !== 200 || resp.data.status !== "ok") {
            throw new Error("#" + resp.data.code + "(" + resp.data.status + ":" + resp.data.message + ")");
        }
        return resp.data.result;
    },
    async change_commission(commission, inviteCodeId, tgUserId) {
        const resp = await axios.post(this._url("change_commission", "commission=" + commission + "&invite_code_id=" + inviteCodeId + "&tg_user_id=" + tgUserId), this.config);
        this._mustOK(resp);
        if (resp.data.code !== 200 || resp.data.status !== "ok") {
            throw new Error("#" + resp.data.code + "(" + resp.data.status + ":" + resp.data.message + ")");
        }
        return resp.data.result;
    },
    async userinfo(tgUserId) {
        const resp = await axios.get(this._url("userinfo", "tg_user_id=" + tgUserId), this.config);
        this._mustOK(resp);
        if (resp.data.code !== 200 || resp.data.status !== "ok") {
            throw new Error("#" + resp.data.code + "(" + resp.data.status + ":" + resp.data.message + ")");
        }
        return resp.data.result;
    },
    async fundlogs(page, pageSize) {
        const resp = await axios.get(this._url("fundlogs"), "page=" + page + "&page_size=" + pageSize, this.config);
        this._mustOK(resp);
        if (resp.data.code !== 200 || resp.data.status !== "ok") {
            throw new Error("#" + resp.data.code + "(" + resp.data.status + ":" + resp.data.message + ")");
        }
        return resp.data.result.items;
    }
};

// 以下为测试代码。
easypay.init("626f63fd5d6816e8087e69d3", "aIjkzXPv3z4sj1VM");
const inspect = function (data) {
    console.log(data)
}

easypay.me().then(inspect);
/*easypay.paylink({
    items: [
        {unique_id:"1",name:"item1",amount:100,auto_renew:true},
        {unique_id:"2",name:"item2",amount:200,auto_renew:false},
        {unique_id:"3",name:"item3",amount:500,auto_renew:false}
    ]
}).then(inspect);*/
// easypay.trans('tron', "12a9100a92a4a4efca1a6f173eb77898a36d15a8df7093e696d175b7dff522a9").then(inspect);
// easypay.order('6275280114fdacffad851af4').then(inspect);
// easypay.orderlist(1, 1000).then(inspect);
/*easypay.transfer({
    order_id: ""+new Date,
    name: "转账",
    amount: 200,
    to_user_id: 928345358
}).then(inspect);*/
/*easypay.deduct({
    order_id: "6275280114fdacffad851af4",
    unique_id: "6275280114fdacffad851af4",
    name: "转账",
    amount: 200,
    use_promo: true
});*/
// easypay.invitecode("ejbu6nfn", "928345358").then(inspect);
// easypay.change_commission(10, "626f881cf0b6d7625ac9c23b", "812342452");
// easypay.userinfo(928345358).then(inspect);
// easypay.fundlogs(1, 1000).then(inspect);
