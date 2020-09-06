package make_proxy

import(
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func MakeProxy(og, target string){
    
    origin, _ := url.Parse(og)
    
    director := func(req *http.Request){
        req.Header.Add("X-Forwarded-Host", req.Host)
        req.Header.Add("X-Origin-Host", origin.Host)
        req.URL.Scheme = "http"
        req.URL.Host = origin.Host
    }
    
    proxy := &httputil.ReverseProxy{Director: director}
    
    http.HandleFunc("/",
        func(w http.ResponseWriter, r *http.Request){
            proxy.ServeHTTP(w,r)
        })
        
    cert := "/home/jc/file-home/go/pems/cert.pem"
    key := "/home/jc/file-home/go/pems/privkey.pem"
    
    log.Fatal(http.ListenAndServeTLS(target, cert, key, nil))

}

// test