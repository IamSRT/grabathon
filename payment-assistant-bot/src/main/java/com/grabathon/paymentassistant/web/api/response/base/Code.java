package com.grabathon.paymentassistant.web.api.response.base;

public enum Code {

    SUCCESS (0, "success"),
    MISSING_REQUIRED_HEADERS (1, "missing mandatory headers"),
    INVALID_LAMBDA_ID (2, "lambda with given id [%s]  does not exists");



    private int code;
    public int getCode () { return code; }

    private String message;
    public String getMessage () { return message; }

    private Code (int code, String message) {
        this.code = code;
        this.message = message;
    }
}
