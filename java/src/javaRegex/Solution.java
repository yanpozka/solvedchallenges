package javaRegex;

import java.util.Scanner;

/**
 *
 * @author yan
 */
public class Solution {

    public static void main(String[] args) {
        
        Scanner in = new Scanner(System.in);
        while (in.hasNext()) {
            String IP = in.next();
            System.out.println(IP.matches(new myRegex().pattern));
        }
    }

    static class myRegex {

        public String pattern;

        public myRegex() {

            String number = "[0-9]{1,2}|[01][0-9]{2}|[2][0-5]{2}";

            pattern = String.format("^(%s)\\.(%s)\\.(%s)\\.(%s)$",
                    number, number, number, number);
        }
    }
}
