package javaInheritence;

/**
 *
 * @author yan
 */
public class Solution {

    public static void main(String[] args) {
        // TODO code application logic here
    }

    abstract class Arithmetic {
        public abstract int add(int a, int b);
    }

    class Adder extends Arithmetic {
        @Override
        public int add(int a, int b) {
            return a + b;
        }
    }
}
