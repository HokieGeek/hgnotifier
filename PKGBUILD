# Maintainer: Andres Perez <andres.f.perez@gmail.com>

pkgname=snotify-git
_pkgname=snotify
epoch=1
pkgver=0
pkgrel=1
pkgdesc='Provides a notification framework'
url='http://github.com/HokieGeek/snotify'
arch=('i686' 'x86_64')
license=('MIT')
source=('git://github.com/HokieGeek/snotify.git')
sha1sums=('SKIP')

provides=("${_pkgname}")
conflicts=("${_pkgname}")

pkgver() {
    cd "${srcdir}/${_pkgname}"
    git log -1 --format='%cd.%h' --date=short | tr -d -
}

build() {
    cd "${srcdir}/${_pkgname}"
    make
}

package() {
    cd "${srcdir}/${_pkgname}"
    prefix=/usr
    make PREFIX="${prefix}" DESTDIR="${pkgdir}" install
    install -Dm644 README.md "${pkgdir}${prefix}/share/doc/${pkgname}/README.md"
}
