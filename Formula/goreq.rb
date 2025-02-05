class Goreq < Formula
    desc "A fast and lightweight HTTP client written in Go"
    homepage "https://github.com/vvvdwbvvv/goreq"
    version "1.0.0"
    url "https://github.com/vvvdwbvvv/goreq/releases/download/v1.0.0/goreq-darwin-arm64.tar.gz"
    sha256 "13595098adcd1be1c8b651fc1394e8ffdbd084fc5d2f60f64cfd9f143f6fd7b3"
  
    def install
      bin.install "goreq"
    end
  
    test do
      assert_match "goreq 1.0.0", shell_output("#{bin}/goreq --version")
    end
  end
  