package utils;

import java.net.InetAddress;
import java.net.UnknownHostException;

public class IpUtils {

    public static String getLocalMachineInfo() {

        String result = "";

        try {
            InetAddress localHost = InetAddress.getLocalHost();

            String hostName = localHost.getHostName();
            String ip = localHost.getHostAddress();

            result = "[" + hostName + "@" + ip + "]";
        } catch (UnknownHostException e) {
            e.printStackTrace();
        }

        return result;
    }

    public static void main(String[] args) {
        System.out.println(getLocalMachineInfo());
    }


}
