package javaHashset;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.HashSet;

/**
 *
 * @author yan
 */
public class Solution {

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int N = Integer.valueOf(br.readLine());
        HashSet<String> set = new HashSet<>();

        while (N-- > 0) {
            set.add(br.readLine());
            System.out.println(set.size());
        }
    }
}
