FROM python:3.12-bookworm
WORKDIR /

RUN rm /etc/localtime && \
    ln -s /usr/share/zoneinfo/Asia/Taipei /etc/localtime

RUN apt update -y && \
    apt install -y libaio1 libnsl2 libc6 curl gcc bash unzip build-essential && \
    apt upgrade -y && \
    cd /tmp && \
    curl -o instantclient-basiclite.zip https://download.oracle.com/otn_software/linux/instantclient/instantclient-basiclite-linuxx64.zip -SL && \
    unzip instantclient-basiclite.zip && \
    mv instantclient*/ /usr/lib/instantclient && \
    rm instantclient-basiclite.zip && \
    ln -s /usr/lib/instantclient/libclntsh.so.21.1 /usr/lib/libclntsh.so && \
    ln -s /usr/lib/instantclient/libocci.so.21.1 /usr/lib/libocci.so && \
    ln -s /usr/lib/instantclient/libociicus.so /usr/lib/libociicus.so && \
    ln -s /usr/lib/instantclient/libnnz21.so /usr/lib/libnnz21.so && \
    ln -s /usr/lib/libnsl.so.2 /usr/lib/libnsl.so.1 && \
    ln -s /lib/libc.so.6 /usr/lib/libresolv.so.2 && \
    ln -s /lib64/ld-linux-x86-64.so.2 /usr/lib/ld-linux-x86-64.so.2 && \
    #apk --no-cache --update add build-base
    apt autoclean && \
    apt clean && \
    apt clean --dry-run

ENV ORACLE_HOME="${PATH}:/usr/lib/instantclient"

RUN python -m pip install oracledb --upgrade --user

RUN pip install --upgrade pip
COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt
COPY . .

CMD ["python", "main.py"]

