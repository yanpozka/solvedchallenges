package javaEndOfFile;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

/**
 *
 * @author yan
 */
public class Solution {

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int c = 1;
        String s;
        while ((s = br.readLine()) != null)
        {
            System.out.printf("%d %s\n", c, s);
            c++; //
        }
    }
}
