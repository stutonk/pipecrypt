PREFIX ?= /usr/local
BINDIR ?= ${PREFIX}/bin
MANDIR ?= ${PREFIX}/share/man

BIN=pipecrypt

${BIN}:
	go build

install: ${BIN}
	install -d ${DESTDIR}${BINDIR}
	install -m 755 ${BIN} ${DESTDIR}${BINDIR}
	install -d ${DESTDIR}${MANDIR}/man1
	install -m 644 ${BIN}.1 ${DESTDIR}${MANDIR}/man1

uninstall:
	rm -f ${DESTDIR}${BINDIR}/${BIN}
	rm -f ${DESTDIR}${MANDIR}/man1/${BIN}.1

clean:
	rm -f ./${BIN}
