<?php

$_easypay_token = "";
$_easypay_errno = -1;
$_easypay_error = "";

function easypay_init($appId, $key) {
    global $_easypay_token = $appId.":".$key;
}

function _easypay_url($method, $queryString=null) {
    global $_easypay_token;
    $extra = "";
    if (isset($queryString)) {
        $extra = "?".$queryString;
    }
	return "https://pay.easypaybot.com/api/merchant/".$_easypay_token."/".$method.$extra;
}

function easypay_do() {
}

function easypay_order($id) {
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, _easypay_url("order", "order_id=".$id));
    // curl_setopt($tuCurl, CURLOPT_POST, 1);
    // curl_setopt($tuCurl, CURLOPT_POSTFIELDS, $data);
    curl_setopt($tuCurl, CURLOPT_HTTPHEADER, array("User-Agent: EasyPayClient/php1.0"));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $output = curl_exec($ch);
    if (curl_errno($ch) !== 0) {
        $_easypay_errno = curl_errno($ch);
        $_easypay_error = curl_error($ch);
        return false;
    }
    curl_close($ch);
    $resp = json_decode($output);
    if (!isset($resp['code']) || !isset($resp['status']) || !isset($resp['result']) || $resp['result'] === false) {
        $_easypay_errno = 65535;
        $_easypay_error = "internal error";
        return false;
    }
    return $resp['result'];
}

function easypay_transfer($params) {
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, _easypay_url("transfer"));
    curl_setopt($tuCurl, CURLOPT_POST, 1);
    $data = json_encode($params);
    curl_setopt($tuCurl, CURLOPT_POSTFIELDS, $data);
    curl_setopt($tuCurl, CURLOPT_HTTPHEADER, array("User-Agent: EasyPayClient/php1.0"));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $output = curl_exec($ch);
    if (curl_errno($ch) !== 0) {
        $_easypay_errno = curl_errno($ch);
        $_easypay_error = curl_error($ch);
        return false;
    }
    curl_close($ch);
    $resp = json_decode($output);
    if (!isset($resp['code']) || !isset($resp['status']) || !isset($resp['result']) || $resp['result'] === false) {
        $_easypay_errno = 65535;
        $_easypay_error = "internal error";
        return false;
    }
    return $resp['result'];
}

function easypay_deduct($params) {
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, _easypay_url("deduct"));
    curl_setopt($tuCurl, CURLOPT_POST, 1);
    $data = json_encode($params);
    curl_setopt($tuCurl, CURLOPT_POSTFIELDS, $data);
    curl_setopt($tuCurl, CURLOPT_HTTPHEADER, array("User-Agent: EasyPayClient/php1.0"));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $output = curl_exec($ch);
    if (curl_errno($ch) !== 0) {
        $_easypay_errno = curl_errno($ch);
        $_easypay_error = curl_error($ch);
        return false;
    }
    curl_close($ch);
    $resp = json_decode($output);
    if (!isset($resp['code']) || !isset($resp['status']) || !isset($resp['result']) || $resp['result'] === false) {
        $_easypay_errno = 65535;
        $_easypay_error = "internal error";
        return false;
    }
    return $resp['result'];
}
