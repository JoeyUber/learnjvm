package jvmgo.book.ch07;
public class InvokeDemo implements Runnable {
  public static void main(String[] args) {
    new InvokeDemo().test();
  }
  public void test() {
    InvokeDemo.staticMethod();              // invokestatic
    InvokeDemo demo = new InvokeDemo();     // invokespecial
    demo.instanceMethod();                  // invokespecial
    boolean e = super.equals(null);                     // invokespecial
	System.out.println(e);
    this.run();                             // invokevirtual
    ((Runnable) demo).run();                // invokeinterface
  }
  public static void staticMethod() {}
  private void instanceMethod() {}
  @Override public void run() {
	  //System.out.println("do run method");
  }
}



