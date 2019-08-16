public class Child extends Pa{
	public void Method(){
		super.Method();
	}
	public static void main(String... args){
		Child c = new Child();
		c.Method();
	}
}