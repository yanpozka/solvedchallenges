package javaStringToken;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Iterator;
import java.util.LinkedList;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 *
 * @author yan
 */
public class Solution {

    /**
     * @param args the command line arguments
     * @throws java.io.IOException
     */
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        Pattern p = Pattern.compile("[A-Za-z]+");
        
        Matcher matcher = p.matcher(br.readLine());
        
        LinkedList<String> list = new LinkedList<>();
        
        while (matcher.find()) {
            list.add(matcher.group());
        }
        
        System.out.println(list.size());
        
        Iterator<String> it = list.iterator();
        
        while (it.hasNext()) {
            System.out.println(it.next());
        }
    }
    
}
