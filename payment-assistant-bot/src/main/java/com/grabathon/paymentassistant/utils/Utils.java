package com.grabathon.paymentassistant.utils;

import java.util.Objects;

import static com.grabathon.paymentassistant.utils.Constants.EMPTY;

public class Utils {

    public static boolean isEmpty (String input) {
        if (Objects.isNull(input)) return true;
        if (EMPTY.equals(input)) return true;
        return false;
    }

}
