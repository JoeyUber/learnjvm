package jvmgo.book.ch03;

import java.util.HashMap;

public class ClassFileTest {
  public static final boolean FLAG = true;
  public static final byte BYTE = 123;
  public static final char X = 'X';
  public static final short SHORT = 12345;
  public static final int INT = 123456789;
  public static final long LONG = 12345678901L;
  public static final float PI = 3.14f;
  public static final double E = 2.71828;
  public static final HashMap A = new HashMap(16);
  public static void main(String[] args) throws RuntimeException {
    System.out.println("Hello, World!");
  }

  public static String staticConvert(Integer val){
    return Integer.toString(val);
  }

  public String convert(Integer val){
    return Integer.toString(val);
  }   
}
