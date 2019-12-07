package com.grabathon.paymentassistant.orchestration.rules.base;

@FunctionalInterface public interface Rule<T> {

    boolean check (T data);

}
