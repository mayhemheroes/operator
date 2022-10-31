package fuzz

import "strconv"
import "github.com/tigera/operator/pkg/render/common/podaffinity"
import "github.com/tigera/operator/pkg/dns"
import "github.com/tigera/operator/pkg/url"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 1 {
        num, _ = strconv.Atoi(string(bytes[0]))

        switch num {
    
        case 0:
            content := string(bytes)
            podaffinity.NewPodAntiAffinity(content, "mayhem")
            return 0

        case 1:
            content := string(bytes)
            dns.GetClusterDomain(content)
            return 0

        default:
            content := string(bytes)
            url.ParseEndpoint(content)
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}