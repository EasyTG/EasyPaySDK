<?php /** @noinspection DuplicatedCode */

$_easypay_token = "";
$_easypay_errno = -1;
$_easypay_error = "";

function easypay_init($appId, $key)
{
    global $_easypay_token;
    $_easypay_token = $appId . ":" . $key;
}

function _easypay_url($method, $queryString = null)
{
    global $_easypay_token;
    $extra = "";
    if (isset($queryString)) {
        $extra = "?" . $queryString;
    }
    $host = 'pay.easypaybot.com';
    // getenv("EASYPAY_DEBUG")
    $debug = getenv("EASYPAY_DEBUG");
    if ($debug !== false && $debug !== "") {
        fprintf(STDERR, "WARNING: DEBUG SERVER USED!!!\n");
        $host = 'pay.test.tgbot.link';
    }
    return "https://$host/api/merchant/" . $_easypay_token . "/" . $method . $extra;
}

function easypay_error()
{
    global $_easypay_errno, $_easypay_error;
    return sprintf("#%d(%s)", $_easypay_errno, $_easypay_error);
}

function easypay_me()
{
    global $_easypay_errno, $_easypay_error;
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, _easypay_url("me"));
    curl_setopt($ch, CURLOPT_HTTPHEADER, array("User-Agent: EasyPaySDK/php1.0", "Content-Type: application/json"));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $output = curl_exec($ch);
    if (curl_errno($ch) !== 0) {
        $_easypay_errno = curl_errno($ch);
        $_easypay_error = curl_error($ch);
        return false;
    }
    curl_close($ch);
    $resp = json_decode($output, true);
    if (!isset($resp['code']) || !isset($resp['status']) || !isset($resp['result']) || $resp['result'] === false) {
        $_easypay_errno = 65535;
        $_easypay_error = "internal error";
        return false;
    }
    return $resp['result'];
}

function easypay_paylink($params)
{
    global $_easypay_errno, $_easypay_error;
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, _easypay_url("paylink"));
    curl_setopt($ch, CURLOPT_HTTPHEADER, array("User-Agent: EasyPaySDK/php1.0", "Content-Type: application/json"));
    curl_setopt($ch, CURLOPT_POST, 1);
    curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($params));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $output = curl_exec($ch);
    if (curl_errno($ch) !== 0) {
        $_easypay_errno = curl_errno($ch);
        $_easypay_error = curl_error($ch);
        return false;
    }
    curl_close($ch);
    $resp = json_decode($output, true);
    if (!isset($resp['code']) || !isset($resp['status']) || !isset($resp['result']['items']) || $resp['result']['items'] === false) {
        $_easypay_errno = 65535;
        $_easypay_error = "internal error";
        return false;
    }
    return $resp['result']['items'];
}

function easypay_trans($type, $txid)
{
    global $_easypay_errno, $_easypay_error;
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, _easypay_url("trans", "type=${type}&txid=${txid}"));
    curl_setopt($ch, CURLOPT_HTTPHEADER, array("User-Agent: EasyPaySDK/php1.0"));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $output = curl_exec($ch);
    if (curl_errno($ch) !== 0) {
        $_easypay_errno = curl_errno($ch);
        $_easypay_error = curl_error($ch);
        return false;
    }
    curl_close($ch);
    $resp = json_decode($output, true);
    if (!isset($resp['code']) || !isset($resp['status']) || !isset($resp['result']) || $resp['result'] === false) {
        $_easypay_errno = 65535;
        $_easypay_error = "internal error";
        return false;
    }
    return $resp['result'];
}

function easypay_order($id)
{
    global $_easypay_errno, $_easypay_error;
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, _easypay_url("order", "order_id=" . $id));
    curl_setopt($ch, CURLOPT_HTTPHEADER, array("User-Agent: EasyPaySDK/php1.0"));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $output = curl_exec($ch);
    if (curl_errno($ch) !== 0) {
        $_easypay_errno = curl_errno($ch);
        $_easypay_error = curl_error($ch);
        return false;
    }
    curl_close($ch);
    $resp = json_decode($output, true);
    if (!isset($resp['code']) || !isset($resp['status']) || !isset($resp['result']) || $resp['result'] === false) {
        $_easypay_errno = 65535;
        $_easypay_error = "internal error";
        return false;
    }
    return $resp['result'];
}


function easypay_orderlist($page, $pageSize)
{
    global $_easypay_errno, $_easypay_error;
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, _easypay_url("orderlist", "page=${page}&page_size=${pageSize}"));
    curl_setopt($ch, CURLOPT_HTTPHEADER, array("User-Agent: EasyPaySDK/php1.0"));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $output = curl_exec($ch);
    if (curl_errno($ch) !== 0) {
        $_easypay_errno = curl_errno($ch);
        $_easypay_error = curl_error($ch);
        return false;
    }
    curl_close($ch);
    $resp = json_decode($output, true);
    if (!isset($resp['code']) || !isset($resp['status']) || !isset($resp['result']['items']) || $resp['result']['items'] === false) {
        $_easypay_errno = 65535;
        $_easypay_error = "internal error";
        return false;
    }
    return $resp['result']['items'];
}

function easypay_transfer($params)
{
    global $_easypay_errno, $_easypay_error;
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, _easypay_url("transfer"));
    curl_setopt($ch, CURLOPT_POST, 1);
    curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($params));
    curl_setopt($ch, CURLOPT_HTTPHEADER, array("User-Agent: EasyPaySDK/php1.0", "Content-Type: application/json"));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $output = curl_exec($ch);
    if (curl_errno($ch) !== 0) {
        $_easypay_errno = curl_errno($ch);
        $_easypay_error = curl_error($ch);
        return false;
    }
    curl_close($ch);
    $resp = json_decode($output, true);
    if (!isset($resp['code']) || !isset($resp['status']) || !isset($resp['result']) || $resp['result'] === false) {
        $_easypay_errno = 65535;
        $_easypay_error = "internal error";
        return false;
    }
    return $resp['result'];
}

function easypay_deduct($params)
{
    global $_easypay_errno, $_easypay_error;
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, _easypay_url("deduct"));
    curl_setopt($ch, CURLOPT_POST, 1);
    curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($params));
    curl_setopt($ch, CURLOPT_HTTPHEADER, array("User-Agent: EasyPaySDK/php1.0", "Content-Type: application/json"));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $output = curl_exec($ch);
    if (curl_errno($ch) !== 0) {
        $_easypay_errno = curl_errno($ch);
        $_easypay_error = curl_error($ch);
        return false;
    }
    curl_close($ch);
    $resp = json_decode($output, true);
    if (!isset($resp['code']) || !isset($resp['status']) || !isset($resp['result']) || $resp['result'] === false) {
        $_easypay_errno = 65535;
        $_easypay_error = "internal error";
        return false;
    }
    return $resp['result'];
}

function easypay_invitecode($code, $tgUserId) {
    global $_easypay_errno, $_easypay_error;
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, _easypay_url("invitecode", "code=${code}&tg_user_id=${tgUserId}"));
    curl_setopt($ch, CURLOPT_HTTPHEADER, array("User-Agent: EasyPaySDK/php1.0"));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $output = curl_exec($ch);
    if (curl_errno($ch) !== 0) {
        $_easypay_errno = curl_errno($ch);
        $_easypay_error = curl_error($ch);
        return false;
    }
    curl_close($ch);
    $resp = json_decode($output, true);
    if (!isset($resp['code']) || !isset($resp['status']) || !isset($resp['result']) || $resp['result'] === false) {
        $_easypay_errno = 65535;
        $_easypay_error = "internal error";
        return false;
    }
    return $resp['result'];
}

function easypay_change_commission($commission, $inviteCodeId, $tgUserId) {
    global $_easypay_errno, $_easypay_error;
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, _easypay_url("change_commission", "commission=${commission}&invite_code_id=${inviteCodeId}&tg_user_id=${tgUserId}"));
    curl_setopt($ch, CURLOPT_HTTPHEADER, array("User-Agent: EasyPaySDK/php1.0", "Content-Type: application/json"));
    curl_setopt($ch, CURLOPT_POST, 1);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $output = curl_exec($ch);
    if (curl_errno($ch) !== 0) {
        $_easypay_errno = curl_errno($ch);
        $_easypay_error = curl_error($ch);
        return false;
    }
    curl_close($ch);
    $resp = json_decode($output, true);
    if (!isset($resp['code']) || !isset($resp['status']) || !isset($resp['result']) || $resp['result'] === false) {
        $_easypay_errno = 65535;
        $_easypay_error = "internal error";
        return false;
    }
    return $resp['result'];
}

function easypay_userinfo($tgUserId) {
    global $_easypay_errno, $_easypay_error;
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, _easypay_url("userinfo", "tg_user_id=${tgUserId}"));
    curl_setopt($ch, CURLOPT_HTTPHEADER, array("User-Agent: EasyPaySDK/php1.0"));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $output = curl_exec($ch);
    if (curl_errno($ch) !== 0) {
        $_easypay_errno = curl_errno($ch);
        $_easypay_error = curl_error($ch);
        return false;
    }
    curl_close($ch);
    $resp = json_decode($output, true);
    if (!isset($resp['code']) || !isset($resp['status']) || !isset($resp['result']) || $resp['result'] === false) {
        $_easypay_errno = 65535;
        $_easypay_error = "internal error";
        return false;
    }
    return $resp['result'];
}

function easypay_fundlogs($page, $pageSize) {
    global $_easypay_errno, $_easypay_error;
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, _easypay_url("fundlogs", "page=${page}&page_size=${pageSize}"));
    curl_setopt($ch, CURLOPT_HTTPHEADER, array("User-Agent: EasyPaySDK/php1.0"));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    $output = curl_exec($ch);
    if (curl_errno($ch) !== 0) {
        $_easypay_errno = curl_errno($ch);
        $_easypay_error = curl_error($ch);
        return false;
    }
    curl_close($ch);
    $resp = json_decode($output, true);
    if (!isset($resp['code']) || !isset($resp['status']) || !isset($resp['result']['items']) || $resp['result']['items'] === false) {
        $_easypay_errno = 65535;
        $_easypay_error = "internal error";
        return false;
    }
    return $resp['result']['items'];
}
