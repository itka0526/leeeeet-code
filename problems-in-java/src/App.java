import java.util.logging.Level;
import java.util.logging.Logger;

public class App
{
    public static final Logger logger = Logger.getLogger("MyLog");

    public static void main(String[] args)
    {
        logger.log(Level.INFO, "");
    }
}
