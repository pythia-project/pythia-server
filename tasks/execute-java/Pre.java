import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.FileOutputStream;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.IOException;
import java.io.OutputStreamWriter;
import java.io.File;
import java.io.FileWriter;
import java.io.Writer;
import java.io.*;
import java.lang.reflect.Field;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.charset.Charset;

public class Pre
{
	private static String ReadProc(InputStream is) throws IOException
	{
		BufferedReader reader = new BufferedReader(new InputStreamReader(is));
			StringBuilder builder = new StringBuilder();
			String line = null;
			while((line = reader.readLine()) != null) {
				builder.append(line);
				builder.append(System.getProperty("line.separator"));
			}
			return  builder.toString();
	}

	public static void main(String[] args)
	{
		try{
			InputStreamReader in = new InputStreamReader(System.in);
			BufferedReader input = new BufferedReader(in);
			String file = "Script.java";
			OutputStreamWriter f = new OutputStreamWriter(new FileOutputStream(file), Charset.forName("UTF-8"));
			String str;

			while((str = input.readLine()) != null) {
				
				//System.out.println(str);
				f.write(str.trim() + "\n");				
			}
			f.close();

			ProcessBuilder pb = new ProcessBuilder( "/usr/lib/jvm/java-8-openjdk-i386/bin/javac","-encoding", "UTF-8", file);
			pb.directory(new File("/tmp/work"));
			Process p = pb.start();
			
			int a = p.waitFor();
			
			if(a != 0) {
				System.out.println("{\"stdout\": \"" + ReadProc(p.getInputStream()) + "\", \"stderr\": \"" + ReadProc(p.getErrorStream()) + "\", \"returncode\": \"" + a + "\"}");
				return;
			}

			pb = new ProcessBuilder( "/usr/lib/jvm/java-8-openjdk-i386/bin/java", "Script");
			p = pb.start();
			
			a = p.waitFor();

			System.out.println("{\"stdout\": \"" + ReadProc(p.getInputStream()) + "\", \"stderr\": \"" + ReadProc(p.getErrorStream()) + "\", \"returncode\": \"" + a + "\"}");

		} catch (IOException|InterruptedException  io) {
			io.printStackTrace();
		}	
	}
}
