import java.util.HashMap;
import java.util.Map;

public class SingleNumberSolution
{
    public int singleNumber(int[] nums)
    {
        HashMap<Integer, Integer> m = new HashMap<>();

        for (int i = 0; i < nums.length; i++)
        {
            if (m.containsKey(nums[i]))
            {
                m.put(nums[i], m.get(nums[i]) + 1);
                continue;
            }
            m.put(nums[i], 1);
        }

        for (Map.Entry<Integer, Integer> e : m.entrySet())
        {
            if (e.getValue() == 1)
            {
                return e.getKey();
            }
        }

        return -1;
    }
}
