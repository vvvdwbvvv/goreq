class Goreq < Formula
    desc "A fast and lightweight HTTP client written in Go"
    homepage "https://github.com/vvvdwbvvv/goreq"
    version "1.0.0"
    url "https://github.com/vvvdwbvvv/goreq/releases/download/v1.0.0/goreq-darwin-arm64.tar.gz"
    sha256 "e76272f9f6e0df22ef53830bdccfaa4442046902302f7694378ec45a8f7e31d2"
  
    def install
      bin.install "goreq"
    end
  
    test do
      assert_match "goreq 1.0.0", shell_output("#{bin}/goreq --version")
    end
  end
  