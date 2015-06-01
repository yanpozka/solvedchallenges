package phoneBook;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.HashMap;

public class Solution {

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        HashMap<String, String> map = new HashMap<>();

        int N = Integer.valueOf(br.readLine());
        while (N-- > 0) {
            String name = br.readLine();
            map.put(name, br.readLine());
        }

        String name;
        while ((name = br.readLine()) != null)
        {
            String num = map.get(name);

            if (num != null) {
                System.out.println(name + "=" + num);
            } else {
                System.out.println("Not found");
            }
        }
    }
}
