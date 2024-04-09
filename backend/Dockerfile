FROM golang:latest

LABEL Maintener(captain)=" Ibrahima GUEYE <https://learn.zone01dakar.sn/git/igueye>"
LABEL Maintener(chat-feature)="Vincent Félix NDOUR <https://learn.zone01dakar.sn/git/vindour> & Madické YADE <https://learn.zone01dakar.sn/git/myade> "
LABEL Maintener(post-feature)="Masseck THIAW <https://learn.zone01dakar.sn/git/mthiaw>"
LABEL Maintener(user-feature)="Mamour Ousmane DRAME <https://learn.zone01dakar.sn/git/mamdrame>"
LABEL Maintener(group-feature)="Ibrahima GUEYE <https://learn.zone01dakar.sn/git/igueye> Nicolas Cor FAYE <https://learn.zone01dakar.sn/git/nifaye>"
LABEL Name="BABBLE SOCIAL"
LABEL service = "back-end"
LABEL version="1.0.0"
LABEL License="BABBLE team"

# working directory 
WORKDIR /backend
# Copy oproject file into container's working dir
COPY . .

EXPOSE 8080
# Command to launch backend service
CMD ["go", "run", "."]
