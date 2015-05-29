package javaapplication1;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;

public class Solution {

    static int nElems = 0;
    static int[] a;

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int N = Integer.valueOf(br.readLine()), num;
        a = new int[N];

        while (N-- > 0) {
            num = Integer.valueOf(br.readLine());
            insert(num);
            // System.out.println(Arrays.toString(a));
            // System.out.println(nElems);

            int len = nElems;
            double r;

            if (len % 2 == 0) {
                r = ((double) a[len / 2] + a[len / 2 - 1]) / 2.0;
            } else {
                r = (double) a[len / 2];
            }
            System.out.println(String.format("%.1f", r));
        }
    }

    public static void insert(int value) {
        int j = Arrays.binarySearch(a, 0, nElems, value);
        if (j < 0) {
            j = -j - 1;
        } else {
            if (a[j] == value) {
                j++;
            }
        }
        System.arraycopy(a, j, a, j + 1, nElems - j);
        a[j] = value;
        nElems++;
    }
}
