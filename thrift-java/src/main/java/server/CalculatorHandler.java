package server;

import gen.share.SharedStruct;
import gen.tutorial.Calculator;
import gen.tutorial.InvalidOperation;
import gen.tutorial.Work;
import org.apache.thrift.TException;
import utils.IpUtils;

import java.util.HashMap;

public class CalculatorHandler implements Calculator.Iface {

    private HashMap<Integer, SharedStruct> log;

    public CalculatorHandler() {
        log = new HashMap<Integer, SharedStruct>();
    }

    public boolean health() throws TException {
        return true;
    }

    public String ping() throws TException {
        System.out.println("pong!");
        return IpUtils.getLocalMachineInfo() + " by Java!";
    }

    public int add(int num1, int num2) throws TException {
        System.out.println("add(" + num1 + "," + num1 + ")");
        return num1 + num2;
    }

    public int calculate(int logid, Work work) throws InvalidOperation, TException {
        System.out.println("calculate(" + logid + ", {" + work.op + "," + work.num1 + "," + work.num2 + "})");
        int val = 0;
        switch (work.op) {
            case ADD:
                val = work.num1 + work.num2;
                break;
            case SUBTRACT:
                val = work.num1 - work.num2;
                break;
            case MULTIPLY:
                val = work.num1 * work.num2;
                break;
            case DIVIDE:
                if (work.num2 == 0) {
                    InvalidOperation io = new InvalidOperation();
                    io.whatOp = work.op.getValue();
                    io.why = "Cannot divide by 0";
                    throw io;
                }
                val = work.num1 / work.num2;
                break;
            default:
                InvalidOperation io = new InvalidOperation();
                io.whatOp = work.op.getValue();
                io.why = "Unknown operation";
                throw io;
        }

        SharedStruct entry = new SharedStruct();
        entry.key = logid;
        entry.value = Integer.toString(val);
        log.put(logid, entry);

        return val;
    }

    public void zip() throws TException {
        System.out.println("zip()");
    }

    public SharedStruct getStruct(int key) throws TException {
        System.out.println("getStruct(" + key + ")");
        return log.get(key);
    }
}
