// go:build !gocv_specific_modules || (gocv_specific_modules && gocv_mcc)

#include "mcc.h"

MccCCheckerDetector MccCCheckerDetector_New()
{
    try
    {
        return new cv::Ptr<cv::mcc::CCheckerDetector>(cv::mcc::CCheckerDetector::create());
    }
    catch (const cv::Exception &e)
    {
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void MccCCheckerDetector_Close(MccCCheckerDetector md)
{
    delete md;
}

MccCCheckerVector MccCCheckerDetector_GetListColorChecker(MccCCheckerDetector md)
{
    try
    {
        std::vector<cv::Ptr<cv::mcc::CChecker>> *result = new std::vector<cv::Ptr<cv::mcc::CChecker>>((*md)->getListColorChecker());
        return result;
    }
    catch (const cv::Exception &e)
    {
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void MccCCheckerVector_Close(MccCCheckerVector mv)
{
    delete mv;
}

int MccCCheckerVector_Size(MccCCheckerVector mv)
{
    return mv->size();
}

MccCChecker MccCCheckerVector_At(MccCCheckerVector mv, int idx)
{
    if (!mv || idx < 0 || idx >= mv->size())
    {
        return NULL;
    }
    return mv->at(idx);
}

MccCChecker MccCCheckerDetector_GetBestColorChecker(MccCCheckerDetector md)
{
    try
    {
        return (*md)->getBestColorChecker();
    }
    catch (const cv::Exception &e)
    {
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

bool MccCCheckerDetector_Process(MccCCheckerDetector md, Mat inputArr, int chartType, int nc, bool useNet)
{
    try
    {
        return (*md)->process(*inputArr, cv::mcc::TYPECHART(chartType), nc);
    }
    catch (const cv::Exception &e)
    {
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

bool MccCCheckerDetector_ProcessWithRegionsOfInterest(MccCCheckerDetector md, Mat inputArr, int chartType, const MccRectVector regionsOfInterest, int nc, bool useNet)
{
    try
    {
        return (*md)->process(*inputArr, cv::mcc::TYPECHART(chartType), *regionsOfInterest, nc, useNet);
    }
    catch (const cv::Exception &e)
    {
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

bool MccCCheckerDetector_SetNet(MccCCheckerDetector md, MccDnnNet net)
{
    try
    {
        return (*md)->setNet(*static_cast<cv::dnn::Net *>(net));
    }
    catch (const cv::Exception &e)
    {
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

Points2f MccCChecker_GetBox(MccCChecker mc)
{
    Points2f out = {nullptr, 0};
    try
    {
        std::vector<cv::Point2f> box = mc->getBox();
        out.length = static_cast<int>(box.size());
        if (out.length > 0)
        {
            out.points = new Point2f[out.length];
            for (int i = 0; i < out.length; ++i)
            {
                out.points[i].x = box[i].x;
                out.points[i].y = box[i].y;
            }
        }
    }
    catch (const cv::Exception &e)
    {
        setExceptionInfo(e.code, e.what());
    }
    return out;
}

Point2f MccCChecker_GetCenter(MccCChecker mc)
{
    try
    {
        cv::Point2f pt = mc->getCenter();
        Point2f out;
        out.x = pt.x;
        out.y = pt.y;
        return out;
    }
    catch (const cv::Exception &e)
    {
        setExceptionInfo(e.code, e.what());
        Point2f out = {0, 0};
        return out;
    }
}

void MccCChecker_SetTarget(MccCChecker mc, int target)
{
    mc->setTarget(cv::mcc::TYPECHART(target));
}
void MccCChecker_SetBox(MccCChecker mc, Point2f *pts, int length)
{
    std::vector<cv::Point2f> box(length);
    for (int i = 0; i < length; ++i)
    {
        box[i] = cv::Point2f(pts[i].x, pts[i].y);
    }
    mc->setBox(box);
}
void MccCChecker_SetChartsRGB(MccCChecker mc, Mat mat)
{
    mc->setChartsRGB(*mat);
}
void MccCChecker_SetChartsYCbCr(MccCChecker mc, Mat mat)
{
    mc->setChartsYCbCr(*mat);
}
void MccCChecker_SetCost(MccCChecker mc, float cost)
{
    mc->setCost(cost);
}
void MccCChecker_SetCenter(MccCChecker mc, Point2f pt)
{
    mc->setCenter(cv::Point2f(pt.x, pt.y));
}

int MccCChecker_GetTarget(MccCChecker mc)
{
    return int(mc->getTarget());
}

Points2f MccCChecker_GetColorCharts(MccCChecker mc)
{
    Points2f out = {nullptr, 0};
    try
    {
        std::vector<cv::Point2f> charts = mc->getColorCharts();
        out.length = static_cast<int>(charts.size());
        if (out.length > 0)
        {
            out.points = new Point2f[out.length];
            for (int i = 0; i < out.length; ++i)
            {
                out.points[i].x = charts[i].x;
                out.points[i].y = charts[i].y;
            }
        }
    }
    catch (const cv::Exception &e)
    {
        setExceptionInfo(e.code, e.what());
    }
    return out;
}

Mat MccCChecker_GetChartsRGB(MccCChecker mc)
{
    return new cv::Mat(mc->getChartsRGB());
}

Mat MccCChecker_GetChartsYCbCr(MccCChecker mc)
{
    return new cv::Mat(mc->getChartsYCbCr());
}

float MccCChecker_GetCost(MccCChecker mc)
{
    return mc->getCost();
}

MccCCheckerDraw MccCCheckerDraw_Create(MccCChecker mc, double b, double g, double r, double a, int thickness)
{
    try
    {
        cv::Scalar color(b, g, r, a);
        return new cv::Ptr<cv::mcc::CCheckerDraw>(cv::mcc::CCheckerDraw::create(cv::Ptr<cv::mcc::CChecker>(mc), color, thickness));
    }
    catch (const cv::Exception &e)
    {
        setExceptionInfo(e.code, e.what());
        return nullptr;
    }
}

void MccCCheckerDraw_Draw(MccCCheckerDraw md, Mat img)
{
    try
    {
        (*md)->draw(*img);
    }
    catch (const cv::Exception &e)
    {
        setExceptionInfo(e.code, e.what());
    }
}

void MccCCheckerDraw_Close(MccCCheckerDraw md)
{
    delete md;
}

MccRectVector MccRectVector_New()
{
    return new std::vector<cv::Rect>();
}

void MccRectVector_PushBack(MccRectVector mv, int x, int y, int width, int height)
{
    mv->emplace_back(x, y, width, height);
}

void MccRectVector_Close(MccRectVector mv)
{
    delete mv;
}

MccDetectorParameters MccDetectorParameters_Create()
{
    try
    {
        return new cv::mcc::DetectorParameters();
    }
    catch (const cv::Exception &e)
    {
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void MccDetectorParameters_Close(MccDetectorParameters mp)
{
    delete mp;
}

void MccDetectorParameters_SetAdaptiveThreshWinSizeMin(MccDetectorParameters mp, int adaptiveThreshWinSizeMin)
{
    mp->adaptiveThreshWinSizeMin = adaptiveThreshWinSizeMin;
}

int MccDetectorParameters_GetAdaptiveThreshWinSizeMin(MccDetectorParameters mp)
{
    return mp->adaptiveThreshWinSizeMin;
}

void MccDetectorParameters_SetAdaptiveThreshWinSizeMax(MccDetectorParameters mp, int adaptiveThreshWinSizeMax)
{
    mp->adaptiveThreshWinSizeMax = adaptiveThreshWinSizeMax;
}

int MccDetectorParameters_GetAdaptiveThreshWinSizeMax(MccDetectorParameters mp)
{
    return mp->adaptiveThreshWinSizeMax;
}

void MccDetectorParameters_SetAdaptiveThreshWinSizeStep(MccDetectorParameters mp, int adaptiveThreshWinSizeStep)
{
    mp->adaptiveThreshWinSizeStep = adaptiveThreshWinSizeStep;
}

int MccDetectorParameters_GetAdaptiveThreshWinSizeStep(MccDetectorParameters mp)
{
    return mp->adaptiveThreshWinSizeStep;
}

void MccDetectorParameters_SetBorderWidth(MccDetectorParameters mp, int borderWidth)
{
    mp->borderWidth = borderWidth;
}

int MccDetectorParameters_GetBorderWidth(MccDetectorParameters mp)
{
    return mp->borderWidth;
}

void MccDetectorParameters_SetMinContourLengthAllowed(MccDetectorParameters mp, int minContourLengthAllowed)
{
    mp->minContourLengthAllowed = minContourLengthAllowed;
}

int MccDetectorParameters_GetMinContourLengthAllowed(MccDetectorParameters mp)
{
    return mp->minContourLengthAllowed;
}

void MccDetectorParameters_SetMinContourPointsAllowed(MccDetectorParameters mp, int minContourPointsAllowed)
{
    mp->minContourPointsAllowed = minContourPointsAllowed;
}

int MccDetectorParameters_GetMinContourPointsAllowed(MccDetectorParameters mp)
{
    return mp->minContourPointsAllowed;
}

void MccDetectorParameters_SetMinImageSize(MccDetectorParameters mp, int minImageSize)
{
    mp->minImageSize = minImageSize;
}

int MccDetectorParameters_GetMinImageSize(MccDetectorParameters mp)
{
    return mp->minImageSize;
}

void MccDetectorParameters_SetMinInterCheckerDistance(MccDetectorParameters mp, int minInterCheckerDistance)
{
    mp->minInterCheckerDistance = minInterCheckerDistance;
}

int MccDetectorParameters_GetMinInterCheckerDistance(MccDetectorParameters mp)
{
    return mp->minInterCheckerDistance;
}

void MccDetectorParameters_SetMinInterContourDistance(MccDetectorParameters mp, int minInterContourDistance)
{
    mp->minInterContourDistance = minInterContourDistance;
}

int MccDetectorParameters_GetMinInterContourDistance(MccDetectorParameters mp)
{
    return mp->minInterContourDistance;
}

void MccDetectorParameters_SetAdaptiveThreshConstant(MccDetectorParameters mp, double adaptiveThreshConstant)
{
    mp->adaptiveThreshConstant = adaptiveThreshConstant;
}

double MccDetectorParameters_GetAdaptiveThreshConstant(MccDetectorParameters mp)
{
    return mp->adaptiveThreshConstant;
}

void MccDetectorParameters_SetConfidenceThreshold(MccDetectorParameters mp, double confidenceThreshold)
{
    mp->confidenceThreshold = confidenceThreshold;
}

double MccDetectorParameters_GetConfidenceThreshold(MccDetectorParameters mp)
{
    return mp->confidenceThreshold;
}

void MccDetectorParameters_SetFindCandidatesApproxPolyDPEpsMultiplier(MccDetectorParameters mp, double findCandidatesApproxPolyDPEpsMultiplier)
{
    mp->findCandidatesApproxPolyDPEpsMultiplier = findCandidatesApproxPolyDPEpsMultiplier;
}

double MccDetectorParameters_GetFindCandidatesApproxPolyDPEpsMultiplier(MccDetectorParameters mp)
{
    return mp->findCandidatesApproxPolyDPEpsMultiplier;
}

void MccDetectorParameters_SetMinContourSolidity(MccDetectorParameters mp, double minContourSolidity)
{
    mp->minContourSolidity = minContourSolidity;
}

double MccDetectorParameters_GetMinContourSolidity(MccDetectorParameters mp)
{
    return mp->minContourSolidity;
}

void MccDetectorParameters_SetMinContoursArea(MccDetectorParameters mp, double minContoursArea)
{
    mp->minContoursArea = minContoursArea;
}

double MccDetectorParameters_GetMinContoursArea(MccDetectorParameters mp)
{
    return mp->minContoursArea;
}

void MccDetectorParameters_SetMinContoursAreaRate(MccDetectorParameters mp, double minContoursAreaRate)
{
    mp->minContoursAreaRate = minContoursAreaRate;
}

double MccDetectorParameters_GetMinContoursAreaRate(MccDetectorParameters mp)
{
    return mp->minContoursAreaRate;
}

void MccDetectorParameters_SetB0factor(MccDetectorParameters mp, float B0factor)
{
    mp->B0factor = B0factor;
}

float MccDetectorParameters_GetB0factor(MccDetectorParameters mp)
{
    return mp->B0factor;
}

void MccDetectorParameters_SetMaxError(MccDetectorParameters mp, float maxError)
{
    mp->maxError = maxError;
}

float MccDetectorParameters_GetMaxError(MccDetectorParameters mp)
{
    return mp->maxError;
}

void MccDetectorParameters_SetMinGroupSize(MccDetectorParameters mp, unsigned minGroupSize)
{
    mp->minGroupSize = minGroupSize;
}

unsigned MccDetectorParameters_GetMinGroupSize(MccDetectorParameters mp)
{
    return mp->minGroupSize;
}
