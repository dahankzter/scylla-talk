package com.scylla.talks;

import com.datastax.driver.core.Cluster;
import com.datastax.driver.core.ProtocolVersion;
import com.datastax.driver.core.Session;

public class Main {

    public static void main(String [] args) {

        Cluster cluster = null;
        Session session = null;
        String node;
        int port;

        try {
            node = args.length >= 1 ? orDefault(args[0], "localhost") : "localhost";
            port = Integer.parseInt(args.length >= 2 ? orDefault(args[1], "9042") : "9042");

            Cluster.Builder b = Cluster.builder().addContactPoint(node);
            b.withPort(port).withProtocolVersion(ProtocolVersion.V4);
            cluster = b.build();
            session = cluster.connect();

            performLogic(session);

        }catch (Exception e) {
            System.out.println(e);
        } finally {
            if (session != null) {
                session.close();
            }
            if (cluster != null) {
                cluster.close();
            }
        }
    }

    private static void performLogic(Session session) {
        session.execute("SELECT id, name, email, about FROM scylla_talk_ks.user").iterator().forEachRemaining(System.out::println);
    }

    private static String orDefault(String val, String def){
        return val != null ? val : def;
    }
}

