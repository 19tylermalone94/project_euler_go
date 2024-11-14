import java.util.stream.IntStream;

public class ThreeOrFive {

    public static void main(String[] args) {
        System.out.println(sumOfMultiplesOf3or5Under1000());
    }

    static int sumOfMultiplesOf3or5Under1000() {
        return IntStream.range(0, 1000)
                        .filter(i -> i % 3 == 0 || i % 5 == 0)
                        .sum();
    }
    
}