<?php
require(__DIR__ . '/sdk.php');

easypay_init("626f63fd5d6816e8087e69d3", "aIjkzXPv3z4sj1VM");
var_dump(easypay_me());
/*var_dump(easypay_paylink([
    'items' => [
        [
            'unique_id' => '1',
            'name' => 'item1',
            'amount' => 100,
            'auto_renew' => true,
        ],
        [
            'unique_id' => '2',
            'name' => 'item2',
            'amount' => 200,
            'auto_renew' => false,
        ],
        [
            'unique_id' => '3',
            'name' => 'item3',
            'amount' => 500,
            'auto_renew' => false,
        ],
    ],
    'params' => 'a=b&c=d',
    'return_url' => 'https://t.me/SuperIndexBot?start=xxxxxxx',
    'expired_at' => time()+86400,
]), easypay_error());*/
// var_dump(easypay_trans("tron", "12a9100a92a4a4efca1a6f173eb77898a36d15a8df7093e696d175b7dff522a9"));
// var_dump(easypay_order('6275280114fdacffad851af4'));
// var_dump(easypay_orderlist(1, 1000));
/*var_dump(easypay_transfer([
    'order_id' => ''.time(),
    'name' => '转账',
    'amount' => 200,
    'to_user_id' => 928345358,
]));*/
/*var_dump(easypay_deduct([
    'order_id' => '6275280114fdacffad851af4',
    'unique_id' => '6275280114fdacffad851af4',
    'name' => '转账',
    'amount' => 200,
    'use_promo' => true,
]));*/
// var_dump(easypay_invitecode('ejbu6nfn', '928345358'));
// var_dump(easypay_change_commission(10, '626f881cf0b6d7625ac9c23b', '812342452'));
// var_dump(easypay_userinfo('928345358'));
// var_dump(easypay_fundlogs(1, 1000));
