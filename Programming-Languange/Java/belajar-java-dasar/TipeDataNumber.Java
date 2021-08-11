public class TipeDataNumber {
    public static void main(String[] args) {
        /*
        ====== Integer ======
        Tipe    | Min                           | Max                       | Size      | Default
        byte    | -128                          | 127                       | 1 byte    | 0
        short   | -32,768                       | 32,767                    | 2 bytes   | 0
        int     | -2,147,483,648                | 2,147,483,647             | 4 bytes   | 0
        long    | -9,223,372,036,854,775,808    | 9,223,372,036,854,775,807 | 8 bytes   | 0

        ====== Desimal ======
        float   | 3.4e−038                      | 3.4e+038                  | 4 bytes   | 0.0
        double  | 1.7e−308                      | 1.7e+308                  | 8 bytes   | 0.0

        */

        byte iniByte = 100;
        short iniShort = 1000;
        int iniInt = 10000000;
        long iniLong1 = 1000000000;
        long iniLong2 = 1000000000L;
        

        System.out.println("iniByte " + iniByte);
        System.out.println("iniShort " + iniShort);
        System.out.println("iniLong1 " + iniLong1);
        System.out.println("iniLong2 " + iniLong2);
    }
}