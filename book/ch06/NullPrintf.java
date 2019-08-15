package book.ch06;
public class NullPrintf{
	public static void main(String... args){
		System.out = null;
		System.out.println("A");
	}
}