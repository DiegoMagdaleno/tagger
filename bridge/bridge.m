//
//  bridge.m
//  tagger
//
//  Created by Diego Magdaleno on 10/30/20.
//

#import <Foundation/Foundation.h>
#include "bridge.h"
#include "ext/NSTaggerURL.h"

const NSArray*
getFilesWithCertainMacOSTag(NSString* path, NSString* targetTag) {
    
    NSMutableArray *fileList = [NSMutableArray array];
    /*
     * Stage 1:
     * We are now on ObjC so we need to convert our Path string to
     * an NSURL
     */
    NSURL *directoryURL = [[NSURL alloc] initFileURLWithPath:path];
    
    NSArray *folders = [[NSFileManager defaultManager] contentsOfDirectoryAtURL:directoryURL includingPropertiesForKeys:NULL options:NSDirectoryEnumerationSkipsHiddenFiles error:NULL];
    
    
    for (id eachPathAsURL in folders) {
        NSURL *target = [NSURL fileURLWithPath:[eachPathAsURL path]];
                
        if ([target GetTags] != nil){
            [fileList addObject:[target path]];
        }
        
        
    }
    return  fileList;
    
}
