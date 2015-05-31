package javaex;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.TreeSet;

public class JavaEx {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        String line = br.readLine();
        
        int N = Integer.valueOf(br.readLine());
        TreeSet<String> set = new TreeSet<>();

        for (int ix = 0, l = line.length(); ix < l - N + 1; ix++) {
            String substring = line.substring(ix, ix+N);
            // System.out.println(substring);
            set.add(substring);
        }
        
        System.out.println(set.first());
        System.out.println(set.last());
    }
}
