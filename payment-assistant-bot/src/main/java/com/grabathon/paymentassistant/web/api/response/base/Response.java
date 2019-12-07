package com.grabathon.paymentassistant.web.api.response.base;

import java.io.Serializable;

public class Response<T> implements Serializable {

    public int code;
    public int getCode() { return code; }

    public String message;
    public String getMessage() { return message; }

    public T data;
    public T getData() { return data; }



    public Response(
            int code,
            String message,
            T data) {

        this.code = code;
        this.message = message;
        this.data = data;
    }

}
