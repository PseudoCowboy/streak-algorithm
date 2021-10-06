import java.awt.List;
import java.util.Map;

public class solution {
    public static void main(String[] args) {
        
    }

    public String destCity(List<List<String>> paths) {
        Map<String, Integer> m = new HashMap<>();
        for (List<String> item : paths) {
            m.computeIfAbsent(item.get(0), v -> 1);
        }
        for (List<String> item : paths) {
            if (m.get(item.get(1)) == 0) {
                return item.get(1);
            }
        }
        return "";
    }
}
