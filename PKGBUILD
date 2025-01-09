# Maintainer: nightvalley totes.maps.2h@icloud.com

pkgname='tammy'
pkgver='1.0.1'
pkgrel=1
pkgdesc='A small CLI utility that will calculate for you the number of lines in all files and directories starting from your current directory.'
arch=('x86_64')
url='https://github.com/nightvalley/tammy'
license=('MIT')
makedepends=('git' 'go')
source=('git+https://github.com/nightvalley/tammy.git')
sha256sums=('SKIP')

pkgver() {
  cd "$srcdir/tammy"
  git describe --tags | sed 's/^v//;s/-/+/g'
}

prepare() {
  cd "$srcdir/tammy"
}

build() {
  cd "$srcdir/tammy"
  go build -o tammy cmd/tammy/main.go
}

package() {
  cd "$srcdir/tammy"
  install -Dm755 tammy "$pkgdir/usr/bin/tammy"
}
