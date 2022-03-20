#!/bin/bash  

myname=$(whoami)

function createDirs() {
    sudo chown $myname:$myname /opt

    mkdir /opt/tools
    mkdir /opt/workspace
    mkdir /opt/project
}


function downloadFile() {
    wget 'https://download.jetbrains.com/idea/ideaIU-2021.3.2.tar.gz' -O /opt/tools/idea.tar.gz
    wget 'https://download.jetbrains.com/go/goland-2021.3.3.tar.gz' -O /opt/tools/goland.tar.gz
    wget 'https://dlcdn.apache.org/maven/maven-3/3.8.5/binaries/apache-maven-3.8.5-bin.tar.gz' -O /opt/tools/maven.tar.gz
    wget 'https://services.gradle.org/distributions/gradle-7.4.1-all.zip' -O /opt/tools/gradle.zip
    wget 'https://go.dev/dl/go1.18.linux-amd64.tar.gz' -O /opt/tools/golang.tar.gz
}

function creatShortCut() {
    tar -xvf /opt/tools/idea.tar.gz -C /opt/tools/
    tar -xvf /opt/tools/goland.tar.gz -C /opt/tools/
    tar -xvf /opt/tools/maven.tar.gz -C /opt/tools/
    tar -xvf /opt/tools/golang.tar.gz -C /opt/tools/
    unzip -d /opt/tools/ /opt/tools/gradle.zip

    ideaVersion=$(find /opt/tools/ -maxdepth 1 -type d -name "*idea*" | awk -F "/" '{print $4}'|uniq)
    golandVersion=$(find /opt/tools/ -maxdepth 1 -type d -name "*GoLand*" | awk -F "/" '{print $4}'|uniq)
    mavenVersion=$(find /opt/tools/ -maxdepth 1 -type d -name "*maven*" | awk -F "/" '{print $4}'|uniq)
    gradleVersion=$(find /opt/tools/ -maxdepth 1 -type d -name "*gradle*" | awk -F "/" '{print $4}'|uniq)

    ideaShortCutPath=/home/$myname/Desktop/idea.desktop
    golandShortCutPath=/home/$myname/Desktop/Goland.desktop

    echo "[Desktop Entry]" >> $ideaShortCutPath
    echo "Version=1.0" >> $ideaShortCutPath
    echo "Type=Application" >> $ideaShortCutPath
    echo "Name=Idea" >> $ideaShortCutPath
    echo "Comment=idea" >> $ideaShortCutPath
    echo "Exec=/opt/tools/$ideaVersion/bin/idea.sh" >> $ideaShortCutPath
    echo "Icon=/opt/tools/$ideaVersion/bin/idea.svg" >> $ideaShortCutPath
    echo "Path=" >> $ideaShortCutPath
    echo "Terminal=false" >> $ideaShortCutPath
    echo "StartupNotify=false" >> $ideaShortCutPath

    echo "[Desktop Entry]" >> $golandShortCutPath
    echo "Version=1.0" >> $golandShortCutPath
    echo "Type=Application" >> $golandShortCutPath
    echo "Name=Goland" >> $golandShortCutPath
    echo "Comment=goland" >> $golandShortCutPath
    echo "Exec=/opt/tools/$golandVersion/bin/goland.sh" >> $golandShortCutPath
    echo "Icon=/opt/tools/$golandVersion/bin/goland.png" >> $golandShortCutPath
    echo "Terminal=false" >> $golandShortCutPath
    echo "StartupNotify=false" >> $golandShortCutPath

    echo "export MAVEN_HOME=\"/opt/tools/$mavenVersion/\"" >> ~/.zshrc
    echo "export GRADLE_HOME=\"/opt/tools/$gradleVersion/\"" >> ~/.zshrc

    echo "export PATH=\"\${PATH}:\$MAVEN_HOME/bin:\$GRADLE_HOME/bin\"" >> ~/.zshrc
}

function banSpyCommunication() {
    echo "0.0.0.0 account.jetbrains.com" >> /etc/hosts
    echo "0.0.0.0 account.jetbrains.com" >> /etc/hosts
    echo "0.0.0.0 www.jetbrains.com" >> /etc/hosts
    echo "0.0.0.0 www-weighted.jetbrains.com" >> /etc/hosts
    echo "0.0.0.0 account.jetbrains.com" >> /etc/hosts
    echo "0.0.0.0 http://www.jetbrains.com" >> /etc/hosts
    echo "0.0.0.0 www-weighted.jetbrains.com" >> /etc/hosts

}

createDirs
downloadFile
creatShortCut
hotfix
banSpyCommunication
