# filesystem_crawler
A simple file system miner in Go


# usage

Takes as input:
  - search root path
  - .txt file path containing one regex per line
  
      ex:<br/>
      
        nano lookup.txt
      Create your regexes <br/>
      
        (?i).txt
        (?i)db
        (?i)coffee

      Then run <br/>

        go run fsys_plun.go /path/to/search/ /path/of/lookup.txt
        
  
      Output:
      
        Parsing :	/****/****/**
        For keys :	[(?i).txt, (?i)db, (?i)coffee]


        [+] Hit
            .name 	LICENSE-MIT.txt
            .perm 	-rw-r--r--
            .@	  /****/****/**/LICENSE-MIT.txt
            .for 	 (?i).txt
            .size	 1068
            .mod 	 2020-12-31 23:50:02.343595328 +0100 CET
            .type	 .txt
        [+] Hit
            .name 	vundle.txt
            .perm 	-rw-r--r--
            .@	    /****/****/**/vundle.txt
            .for 	  (?i).txt
            .size	  14854
            .mod 	  2020-12-31 23:50:02.34590733 +0100 CET
            .type	  .txt
        [+] Hit
            .name 	pwd.txt
            .perm 	-rw-r--r--
            .@	    /****/****/**coffee.csv
            .for 	  (?i).txt
            .size	  7
            .mod 	  2021-01-01 02:27:56.19775488 +0100 CET
            .type	  .txt

