package com.grabathon.paymentassistant.orchestration.rules;

import com.google.gson.Gson;
import com.grabathon.paymentassistant.http.HTTP;
import org.json.JSONObject;

public class Lambdas {
     private static final String walletServiceHost  = "http://192.168.43.141:4000/api/";
     private static final String Success  = "Success";

     public static Float GetBalance(JSONObject stepsRequest){
         String url = walletServiceHost + "user/" +stepsRequest.getString("UserId");
         JSONObject response = HTTP.get(url);
         if (response == null){return (float) -1;}
         return  response.getJSONObject("data").getJSONObject("wallet").getFloat("balance");
     }

    public static String GetUser(JSONObject stepsRequest){
        String url = walletServiceHost + "user/" +stepsRequest.getString("UserId");
        JSONObject response = HTTP.get(url);
        if (response == null){return "{}";}
        return  response.toString();
    }

    public static boolean requestVouch (JSONObject stepsRequest){
        String url = walletServiceHost + "user/vouch";
        String data = new Gson().toJson(stepsRequest);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public static boolean acceptVouch(JSONObject stepsRequest){
        String url = walletServiceHost + "user/vouch" + stepsRequest.getString("vouch_id") + "/accept";
        String data = new Gson().toJson(stepsRequest);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public static boolean declineVouch(JSONObject stepsRequest){
        String url = walletServiceHost + "user/vouch/" + stepsRequest.getString("vouch_id") + "/reject";
        String data = new Gson().toJson(stepsRequest);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public static boolean EnableAutoPay(JSONObject stepsRequest){
        String url = walletServiceHost + "user/enable-auto_pay";
        String data = new Gson().toJson(stepsRequest);
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public static boolean IsAutoPayEnabled(JSONObject stepsRequest){
        String url = walletServiceHost + "user/is-auto-pay-enabled";
        String data = new Gson().toJson(stepsRequest.getJSONObject("Vouch"));
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public static boolean IsVouchValid(JSONObject stepsRequest){
        String url = walletServiceHost + "user/is-vouch-valid";
        String data = new Gson().toJson(stepsRequest.getJSONObject("Vouch"));
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public static boolean CreatePayment(JSONObject stepsRequest){
        String url = walletServiceHost + "payment/create";
        String data = new Gson().toJson(stepsRequest.getJSONObject("payment"));
        JSONObject response = HTTP.post(url,data);
        if (response == null){return false;}
        return response.getString("message").equals(Success);
    }

    public static String GetPayment(JSONObject stepsRequest){
        String url = walletServiceHost + "payment/"+stepsRequest.get("PaymentId");
        JSONObject response = HTTP.get(url);
        if (response == null){return "{}";}
        return response.toString();
    }

    public static JSONObject GetPendingPayment(JSONObject stepsRequest){
        String url = walletServiceHost + "user/"+stepsRequest.get("UserId")+"/pending-payments";
        JSONObject response = HTTP.get(url);
        return response;
    }
}
