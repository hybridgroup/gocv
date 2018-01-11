#include "face_recognizer.h"

// FaceRecognizer
FaceRecognizer FaceRecognizer_New() 
{
    return new cv::Ptr<cv::pvl::FaceRecognizer>(cv::pvl::FaceRecognizer::create());
}

void FaceRecognizer_Close(FaceRecognizer f) 
{
    delete f;
}

void FaceRecognizer_Clear(FaceRecognizer f) {
    (*f)->clear();
}

bool FaceRecognizer_Empty(FaceRecognizer f) {
    return (*f)->empty();
}

void FaceRecognizer_SetTrackingModeEnabled(FaceRecognizer f, bool enabled)
{
    (*f)->setTrackingModeEnabled(enabled);
    return;
}

int FaceRecognizer_GetNumRegisteredPersons(FaceRecognizer f) {
    return (*f)->getNumRegisteredPersons();
}

int FaceRecognizer_CreateNewPersonID(FaceRecognizer f) {
    return (*f)->createNewPersonID();
}

void FaceRecognizer_Recognize(FaceRecognizer f, Mat img, Faces faces, IntVector pids, IntVector confs) {
    std::vector<cv::pvl::Face> vFaces;
    for (size_t i = 0; i < faces.length; ++i) {
        vFaces.push_back(cv::pvl::Face(*faces.faces[i]));
    }

    std::vector<int> personIDs;
    std::vector<int> confidence;

    (*f)->recognize(*img, vFaces, personIDs, confidence);

    pids.val = &personIDs[0];
    pids.length = personIDs.size();

    confs.val = &confidence[0];
    confs.length = confidence.size();

    return;
}

int64_t FaceRecognizer_RegisterFace(FaceRecognizer f, Mat img, Face face, int personID, bool saveTofile) {
    return (*f)->registerFace(*img, cv::pvl::Face(*face), personID, saveTofile);
}

void FaceRecognizer_DeregisterFace(FaceRecognizer f, int64_t faceID) {
    (*f)->deregisterFace(faceID);
}

void FaceRecognizer_DeregisterPerson(FaceRecognizer f, int personID) {
    (*f)->deregisterPerson(personID);
}

void FaceRecognizer_Read(FaceRecognizer f, const char* filename) {
    cv::FileStorage fsRead(filename, cv::FileStorage::READ);
    (*f)->read(fsRead.root());
}

void FaceRecognizer_Save(FaceRecognizer f, const char* filename) {
    (*f)->save(filename);
}
