package javaLoops;

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

        int T = Integer.valueOf(br.readLine());

        while (T-- > 0) {
            String[] parts = br.readLine().split(" ");
            int a = Integer.valueOf(parts[0]),
                    b = Integer.valueOf(parts[1]),
                    N = Integer.valueOf(parts[2]);
            
            for (int ix = 0; ix < N; ix++) {
                int sum = 0;
                for (int k = 0; k <= ix; k++) {
                    sum += (b * Math.pow(2, k));
                }
                sum += a;
                System.out.printf("%d ", sum);
            }
            System.out.println();
        }
    }
}
